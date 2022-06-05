package v1

import (
	context "context"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterGoodsHttpServer(engine *gin.Engine, server GoodsServiceServer) {
	g := engine.Group(Prefix())
	// g := engine.Group("/api/goods/v1/")
	{
		g.GET("/find/:id", FindGoodsTransfer(server.FindGoods))
		g.POST("/new", NewGoodsTransfer(server.NewGoods))
		g.POST("/sale", SaleGoodsTransfer(server.SaleGoods))
		g.DELETE("/delete/:id", DeleteGoodsTransfer(server.DeleteGoods))
	}
}

// get api prefix according file location
func Prefix() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to assign router prefix")
	}
	dirFrag := strings.Split(strings.Replace(file, dir, "", -1), "/")
	dirFrag = dirFrag[:len(dirFrag)-1]
	return strings.Join(dirFrag, "/")
}

// transfer restful request method to gin handler functions
func FindGoodsTransfer(f func(ctx context.Context, in *FindGoodsRequest) (*GoodsReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		in := new(FindGoodsRequest)
		// id, err := strconv.Atoi(c.Param("id"))
		id, err := strconv.ParseInt((c.Param("id")), 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		in.Id = id
		goods, err := f(context.Background(), in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, goods)
	}
}

func SaleGoodsTransfer(f func(ctx context.Context, in *SaleGoodsRequest) (*GoodsReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in SaleGoodsRequest
		if err := c.ShouldBind(&in); err != nil {
			c.String(http.StatusBadRequest, "missing field")
			return
		}
		goods, err := f(context.Background(), &in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, goods)
	}
}

func DeleteGoodsTransfer(f func(ctx context.Context, in *DeleteGoodsRequest) (*DeleteGoodsReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		in := new(DeleteGoodsRequest)
		id, err := strconv.ParseInt((c.Param("id")), 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		in.Id = id
		goods, err := f(context.Background(), in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, goods)
	}
}

func NewGoodsTransfer(f func(ctx context.Context, in *NewGoodsRequest) (*GoodsReply, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in NewGoodsRequest
		if err := c.ShouldBind(&in); err != nil {
			c.String(http.StatusBadRequest, "missing field")
			return
		}
		goods, err := f(context.Background(), &in)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.AsciiJSON(http.StatusOK, goods)
	}
}
