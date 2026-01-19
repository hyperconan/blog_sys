package types

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreatePostResponse struct {
	PostId uint32 `json:"post_id"`
}

type GetPostsResponse struct {
	Posts []PostInfo `json:"posts"`
}

type PostInfo struct {
	Id       uint32 `json:"id"`
	Title    string `json:"title"`
	UserId   uint32 `json:"user_id"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type UpdatePostRequest struct {
	PostId  uint32 `path:"post_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostResponse struct {
	Message string `json:"message"`
}

type DeletePostRequest struct {
	PostId uint32 `path:"post_id"`
}

type DeletePostResponse struct {
	Message string `json:"message"`
}

type CreateCommentRequest struct {
	PostId  uint32 `json:"post_id"`
	Content string `json:"content"`
}

type CreateCommentResponse struct {
	CommentId uint32 `json:"comment_id"`
}

type GetCommentsRequest struct {
	PostId uint32 `path:"post_id"`
}

type GetCommentsResponse struct {
	Comments []CommentInfo `json:"comments"`
}

type CommentInfo struct {
	Id       uint32 `json:"id"`
	Content  string `json:"content"`
	UserId   uint32 `json:"user_id"`
	PostId   uint32 `json:"post_id"`
	CreateAt string `json:"create_at"`
}