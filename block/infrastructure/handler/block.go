package handler

import (
	"fmt"
	"log"
	"strings"

	d "github.com/bruno5200/CSM/block/domain"
	"github.com/bruno5200/CSM/block/infrastructure/client"
	p "github.com/bruno5200/CSM/block/infrastructure/presenter"
	u "github.com/bruno5200/CSM/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *blockHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		log.Printf("Error parsing id: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(d.ErrInvalidBlockId))
	}

	block, err := h.BlockService.GetBlock(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(p.BlockErrorResponse(d.ErrGettingBlock))
	}

	if err := u.CheckDir(d.FilesDir); err != nil {
		return err
	}

	filePath := d.FilesDir + block.Name

	if err := client.NewClient().DownloadFromBlobStorage(fmt.Sprintf("%s/%s.%s", blobPath(block.Extension), block.Id, block.Extension), filePath); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	return c.Download(filePath, block.Name)
}

func (h *blockHandler) GetParam(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		log.Printf("Error parsing id: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(p.BlockErrorResponse(err))
	}

	block, err := h.BlockService.GetBlock(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(p.BlockErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(p.BlockSuccessResponse(block))
}

func blobPath(extension string) string {

	switch strings.ToLower(strings.ReplaceAll(extension, ".", "")) {
	case "pdf", "doc", "docx", "xls", "xlsx", "ppt", "pptx", "md", "html", "csv", "xml", "json", "yaml", "yml", "toml", "txt":
		return e.GetBlobUrl("documents")
	case "jpg", "jpeg", "png", "gif", "svg", "webp", "bmp", "ico", "tiff", "tif":
		return e.GetBlobUrl("images")
	default:
		return e.GetBlobUrl("")
	}
}
