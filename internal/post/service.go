package post

type PostService interface {
}

type postServiceImpl struct {
}

var _ PostService = (*postServiceImpl)(nil)

var Service = postServiceImpl{}
