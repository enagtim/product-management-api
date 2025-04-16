package category

type CategoryService struct {
	CategoryRepository *CategoryRepository
}

func NewCategoryService(categoryRepository *CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepository: categoryRepository,
	}
}

func (s *CategoryService) Create(payload *CategoryPayloadRequest) (*Category, error) {
	category := Category{Name: payload.Name}
	return s.CategoryRepository.Create(&category)
}
