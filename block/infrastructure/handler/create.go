package handler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	d "github.com/bruno5200/CSM/block/domain"
	"github.com/bruno5200/CSM/block/infrastructure/client"
	p "github.com/bruno5200/CSM/block/infrastructure/presenter"
	u "github.com/bruno5200/CSM/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils/v2"
	"github.com/google/uuid"
)

func (h *blockHandler) Post(c *fiber.Ctx) error {

	serviceId, err := uuid.Parse(utils.CopyString(c.Get(d.HeaderServiceId)))

	if err != nil {
		log.Printf("Error parsing serviceId: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(d.ErrInvalidServiceId))
	}

	groupId, err := uuid.Parse(c.Params("groupId"))

	if err != nil {
		log.Printf("Error parsing groupId: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(d.ErrInvalidGroupId))
	}

	file, err := c.FormFile("file")

	if err != nil {
		log.Printf("Error getting file: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(d.ErrMalformedFormKey))
	}

	if !strings.Contains(file.Filename, ".") {
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(d.ErrInvalidFileExtension))
	}

	if err := u.CheckDir(d.FilesDir); err != nil {
		return err
	}

	id := uuid.New()

	ext := strings.ToLower(filepath.Ext(file.Filename))

	filePath := fmt.Sprintf("%s/%s%s", d.FilesDir, id, ext)

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		if err := os.Remove(filePath); err != nil {
			log.Printf("Error removing file: %s", err)
			return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
		}
	}

	if err := c.SaveFile(file, filePath); err != nil {
		log.Printf("Error saving file: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	checksum, err := u.CalculateSHA256Checksum(filePath)

	if err != nil {
		log.Printf("Error calculating checksum: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	log.Printf("Checksum: %s", checksum)

	if block, err := h.BlockService.GetBlockByCheksum(checksum); err == nil {
		return c.Status(fiber.StatusAccepted).JSON(p.BlockCreateResponse(fmt.Sprintf("%s/%s", e.GetUrl(), block.Id)))
	} else {
		log.Printf("DB: %s", err)
	}

	var url, key string

	switch ext {
	case ".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".md", ".html", ".csv", ".xml", ".json", ".yaml", ".yml", ".toml", ".txt":
		url = e.GetBlobUrl("documents")
		key = "document"
	case ".jpg", ".jpeg", ".png", ".gif", ".svg", ".webp", ".bmp", ".ico", ".tiff", ".tif":
		url = e.GetBlobUrl("images")
		key = "image"
	default:
		url = e.GetBlobUrl("")
		key = "document"
	}

	block := d.NewBlock(file.Filename, checksum, fmt.Sprintf("%s/%s", url, id), strings.ToUpper(strings.ReplaceAll(ext, ".", "")), id, groupId, serviceId)

	filechan := make(chan string, 1)

	go u.FileToBase64(file, filechan)

	b64string := <-filechan

	log.Printf("Base64 File: %s...", b64string[:25])

	if err := client.NewClient().UploadToBlob(b64string, url, key, fmt.Sprintf("%s%s", id, ext)); err != nil {
		log.Printf("Error uploading to blob: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	location := fmt.Sprintf("%s/%s", e.GetUrl(), block.Id)

	if err := h.BlockService.CreateBlock(block); err != nil {
		log.Printf("DB: %s", err)
		return c.Status(fiber.StatusNotFound).JSON(p.BlockErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(p.BlockCreateResponse(location))
}

func (h *blockHandler) Json(c *fiber.Ctx) error {

	serviceId, err := uuid.Parse(utils.CopyString(c.Get(d.HeaderServiceId)))

	if err != nil {
		log.Printf("Error parsing serviceId: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(d.ErrInvalidServiceId))
	}

	groupId, err := uuid.Parse(c.Params("groupId"))

	if err != nil {
		log.Printf("Error parsing groupId: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(d.ErrInvalidGroupId))
	}

	blockRequest, err := d.UnmarshalBlockRequest(c.Body())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	if strings.Contains(blockRequest.Extension, ".") {
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(d.ErrInvalidFileExtension))
	}

	fileBytes, err := u.StringBase64ToBytes(blockRequest.Content)

	if err != nil {
		log.Printf("Error converting base64 to bytes: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	checksum, err := u.CalculateSHA256ChecksumBytes(fileBytes)

	if err != nil {
		log.Printf("Error calculating checksum: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	log.Printf("Checksum: %s", checksum)

	if block, err := h.BlockService.GetBlockByCheksum(checksum); err == nil {
		log.Printf("File already exists: %s", block.Id)
		return c.Status(fiber.StatusAccepted).JSON(p.BlockCreateResponse(fmt.Sprintf("%s/%s", e.GetUrl(), block.Id)))
	}

	ext := strings.ToLower(blockRequest.Extension)

	var url, key string

	switch ext {
	case "pdf", "doc", "docx", "xls", "xlsx", "ppt", "pptx", "md", "html", "csv", "xml", "json", "yaml", "yml", "toml", "txt":
		url = e.GetBlobUrl("documents")
		key = "document"
	case "jpg", "jpeg", "png", "gif", "svg", "webp", "bmp", "ico", "tiff", "tif":
		url = e.GetBlobUrl("images")
		key = "image"
	default:
		url = e.GetBlobUrl("")
		key = "document"
	}

	id := uuid.New()

	block := d.NewBlock(fmt.Sprintf("%s.%s", blockRequest.Name, strings.ToLower(blockRequest.Extension)), checksum, fmt.Sprintf("%s/%s", url, id), strings.ToUpper(blockRequest.Extension), id, groupId, serviceId)

	location := fmt.Sprintf("%s/%s", e.GetUrl(), block.Id)

	if err := client.NewClient().UploadToBlob(blockRequest.Content, url, key, fmt.Sprintf("%s%s", id, ext)); err != nil {
		log.Printf("Error uploading to blob: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	if err := h.BlockService.CreateBlock(block); err != nil {
		log.Printf("DB: %s", err)
		return c.Status(fiber.StatusNotFound).JSON(p.BlockErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(p.BlockCreateResponse(location))
}
