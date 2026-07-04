package internal

type createPostRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type updatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type addCommentRequest struct {
	AuthorID int    `json:"author_id"`
	Content  string `json:"content" validate:"required"`
}

type registerRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

type ApiResponseModel struct {
	Message  string `json:"message"`
	Data     any    `json:"data,omitzero"`
	Metadata any    `json:"metadata,omitzero"`
	Error    any    `json:"error,omitzero"`
}

type getPostDetailResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorName string `json:"author_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type getPostListModel struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type getCommentListModel struct {
	ID         int    `json:"id"`
	PostID     int    `json:"post_id"`
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	Content    string `json:"content"`

	CreatedAt string `json:"created_at"`
}
