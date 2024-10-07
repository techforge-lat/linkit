package user

type Handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h Handler) Create() error {
	if err := h.useCase.Create(&User{}); err != nil {
		return err
	}

	return nil
}
