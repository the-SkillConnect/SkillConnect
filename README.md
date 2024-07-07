# SkillConnect: A Freelancing Telegram Mini App (TMA)

Welcome to SkillConnect! This project aims to create a Telegram mini app for freelancers. Follow these steps to get started:

## Getting the Project

1. Clone the repository:

```bash
git clone --recursive https://github.com/the-SkillConnect/SkillConnect.git

cd SkillConnect
```

## Running the Project

1. Set up the database (PostgreSQL):

```bash
sudo docker compose up -d
```

- > [!NOTE] You can populate the database with sample data using `go run scripts/seed.go`

- > [!NOTE] If you encounter issues downloading dependency packages, set the Go proxy using: `go env -w GOPROXY=https://goproxy.io,direct`

2. Run the SkillConnect program:

```bash
go run .
```

## Contributing

We welcome contributions! If you’d like to contribute to SkillConnect, follow these steps:

1. Fork the repository on GitHub.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your branch to your forked repository.
5. Open a pull request to merge your changes into the main repository.

Feel free to customize the contribution guidelines further based on your project’s specific needs. Happy coding!
