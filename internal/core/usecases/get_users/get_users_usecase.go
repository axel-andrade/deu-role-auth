package get_users

type GetUsersUC struct {
	Gateway GetUsersGateway
}

func BuildGetUsersUC(g GetUsersGateway) *GetUsersUC {
	return &GetUsersUC{Gateway: g}
}

func (bs *GetUsersUC) Execute(input GetUsersInputDTO) (*GetUsersOutputDTO, error) {
	users, totalUsers, err := bs.Gateway.GetUsersPaginate(input.PaginationOptions)
	if err != nil {
		return nil, err
	}

	return &GetUsersOutputDTO{users, totalUsers}, nil
}
