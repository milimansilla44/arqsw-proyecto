package services

import (
	productCliente "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/clients/product"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/dto"
	e "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/errors"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/model"

	"strings"
)

type productService struct{}

type productServiceInterface interface {
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
	GetProductsByPalabrasClaves(clave string) (dto.ProductsDto, e.ApiError)
	GetProductsByCategory(category int) (dto.ProductsDto, e.ApiError)
	GetCategories() ([]dto.CategoryDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}
func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {

	var product model.Product = productCliente.GetProductById(id)
	var productDto dto.ProductDto

	if product.Id == 0 {
		return productDto, e.NewBadRequestApiError("product not found")
	}

	productDto.Id = product.Id
	productDto.Name = product.Name
	productDto.Description = product.Description
	productDto.Picture = product.Picture
	productDto.Price = product.Price
	productDto.Stock = product.Stock
	productDto.CategoryId = product.CategoryId

	return productDto, nil
}

func (s *productService) GetProducts() (dto.ProductsDto, e.ApiError) {
	var products model.Products = productCliente.GetProducts()
	var productsDto dto.ProductsDto

	for _, product := range products {

		var productDto dto.ProductDto

		productDto.Id = product.Id
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Picture = product.Picture
		productDto.Price = product.Price
		productDto.Stock = product.Stock
		productDto.CategoryId = product.CategoryId

		productsDto = append(productsDto, productDto)
	}
	return productsDto, nil
}

func (s *productService) GetProductsByPalabrasClaves(clave string) (dto.ProductsDto, e.ApiError) {
	var products model.Products = productCliente.GetProducts()
	var productsDto dto.ProductsDto
	var control int = 0
	for _, product := range products {
		var aux string = strings.ToLower(product.Name)
		var aux2 string = strings.ToLower(product.Description)
		if strings.Contains(aux, clave) || strings.Contains(aux2, clave) {
			control = 1

			var productDto dto.ProductDto

			productDto.Id = product.Id
			productDto.Name = product.Name
			productDto.Description = product.Description
			productDto.Picture = product.Picture
			productDto.Price = product.Price
			productDto.Stock = product.Stock
			productDto.CategoryId = product.CategoryId

			productsDto = append(productsDto, productDto)
		}

	}
	if control == 0 {
		return productsDto, e.NewBadRequestApiError("products not found")
	}

	return productsDto, nil
}

func (s *productService) GetProductsByCategory(category int) (dto.ProductsDto, e.ApiError) {
	var products model.Products = productCliente.GetProducts()
	var productsDto dto.ProductsDto

	for _, product := range products {

		if category == product.CategoryId {

			var productDto dto.ProductDto

			productDto.Id = product.Id
			productDto.Name = product.Name
			productDto.Description = product.Description
			productDto.Picture = product.Picture
			productDto.Price = product.Price
			productDto.Stock = product.Stock
			productDto.CategoryId = product.CategoryId

			productsDto = append(productsDto, productDto)
		}
	}
	return productsDto, nil
}

func (s *productService) GetCategories() ([]dto.CategoryDto, e.ApiError) {
	var categories []model.Category = productCliente.GetCategories()
	var categoriesDTO []dto.CategoryDto

	for _, category := range categories {

		var categoryDTO dto.CategoryDto

		categoryDTO.Id = category.Id
		categoryDTO.Name = category.Name

		categoriesDTO = append(categoriesDTO, categoryDTO)
	}
	return categoriesDTO, nil
}
