package backofficeuserservice



type repository interface {
	// ListAllUsers() ([]entity.User, error)
	// DeleteUser(UserID uint)(error)
	// ListAllLaptops() ([]entity.Laptop, error)
	// DeleteLaptop(LaptopID uint)(error)
	// UpdateLaptop(updatedLaptop entity.Laptop) (error)
	
}

type Service struct {
	repo repository
}

func New(repo repository) Service {
	return Service{repo: repo}
}

