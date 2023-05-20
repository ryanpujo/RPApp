package controller

import (
	"context"

	"github.com/ryanpujo/product-service/internal/pserror"
	"github.com/ryanpujo/product-service/internal/repository"
	"github.com/ryanpujo/product-service/internal/service"
	"github.com/ryanpujo/product-service/product-proto/grpc/product"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type productController struct {
	product.UnimplementedProductServiceServer
	productService service.ProductService
}

func NewProductController(ps service.ProductService) *productController {
	return &productController{productService: ps}
}

func (pc productController) CreateProduct(ctx context.Context, payload *product.ProductPayload) (*product.CreatedProduct, error) {
	args := repository.CreateProductParams{
		StoreID:     payload.StoreId,
		Name:        payload.GetName(),
		Description: payload.GetDescription(),
		Price:       payload.GetPrice(),
		ImageUrl:    payload.GetImageUrl(),
		Stock:       int32(payload.GetStock()),
		CategoryID:  payload.GetCategoryId(),
	}
	createdProduct, err := pc.productService.CreateProduct(ctx, args)

	if err != nil {
		return nil, pserror.ToGrpcError(err)
	}

	product := &product.CreatedProduct{
		Id:          createdProduct.ID,
		StoreId:     createdProduct.StoreID,
		Name:        createdProduct.Name,
		Description: createdProduct.Description,
		Price:       createdProduct.Price,
		ImageUrl:    createdProduct.ImageUrl,
		Stock:       uint32(createdProduct.Stock),
		CategoryId:  int64(createdProduct.CategoryID),
		CreatedAt:   timestamppb.New(createdProduct.CreatedAt.Time),
	}

	return product, nil
}

func (pc productController) DeleteProduct(ctx context.Context, id *product.ProductID) (*emptypb.Empty, error) {
	err := pc.productService.DeleteProduct(ctx, id.GetId())

	return &emptypb.Empty{}, pserror.ToGrpcError(err)
}

func (pc productController) GetProductById(ctx context.Context, id *product.ProductID) (*product.Product, error) {
	result, err := pc.productService.GetProductByID(ctx, id.GetId())

	if err != nil {
		return nil, pserror.ToGrpcError(err)
	}

	foundProduct := product.Product{
		Id:          result.ID,
		StoreName:   result.StoreName,
		Name:        result.Name,
		Description: result.Description,
		Price:       result.Price,
		ImageUrl:    result.ImageUrl,
		Stock:       uint32(result.Stock),
		Category:    result.Category,
		CreatedAt:   timestamppb.New(result.CreatedAt.Time),
	}
	return &foundProduct, nil
}

func (pc productController) GetProducts(ctx context.Context, emptypb *emptypb.Empty) (*product.Products, error) {
	results, err := pc.productService.GetProducts(ctx)
	products := product.Products{
		Products: make([]*product.Product, 0, len(results)),
	}

	if err != nil {
		return nil, pserror.ToGrpcError(err)
	}

	for _, v := range results {
		product := product.Product{
			Id:          v.ID,
			StoreName:   v.StoreName,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			ImageUrl:    v.ImageUrl,
			Stock:       uint32(v.Stock),
			Category:    v.Category,
			CreatedAt:   timestamppb.New(v.CreatedAt.Time),
		}
		products.Products = append(products.Products, &product)
	}

	return &products, nil
}

func (pc productController) UpdateProduct(ctx context.Context, payload *product.ProductPayload) (*emptypb.Empty, error) {
	args := repository.UpdateProductParams{
		ID:          payload.Id,
		StoreID:     payload.StoreId,
		Name:        payload.GetName(),
		Description: payload.GetDescription(),
		Price:       payload.GetPrice(),
		ImageUrl:    payload.GetImageUrl(),
		Stock:       int32(payload.GetStock()),
		CategoryID:  payload.GetCategoryId(),
	}

	err := pc.productService.UpdateProduct(ctx, args)

	return &emptypb.Empty{}, pserror.ToGrpcError(err)
}
