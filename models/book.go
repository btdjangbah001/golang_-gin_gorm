package models

type Book struct {
	ID     uint   `json:"id" gorm: "primary-key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FindAllBooks() (*[]Book, error) {
	var books []Book

	err := DB.Find(&books).Error

	if err != nil {
		return &[]Book{}, err
	}

	return &books, nil
}

func (b *Book) SaveBook() (*Book, error) {
	err := DB.Create(&b).Error

	if err != nil {
		return &Book{}, err
	}

	return b, nil
}

func FindBook(id string) (*Book, error) {
	var book Book
	err := DB.Where("id = ?", id).First(&book).Error

	if err != nil {
		return &Book{}, err
	}
	return &book, nil
}

func (b *Book) UpdateBook(input *UpdateBookInput) (*Book, error) {
	// input.Author = strings.TrimSpace(input.Author)
	// input.Title = strings.TrimSpace(input.Title)

	// if len(input.Author) > 0 {
	// 	b.Author = input.Author
	// }
	// if len(input.Title) > 0 {
	// 	b.Title = input.Title
	// }

	// err := DB.Save(&b).Error

	err := DB.Model(*b).Updates(*input).Error

	if err != nil {
		return &Book{}, err
	}
	return b, nil
}

func (b *Book) DeleteBook() (*Book, error) {
	err := DB.Delete(b).Error

	if err != nil {
		return b, err
	}

	return b, nil
}
