<h1 align="center">Welcome to sygnus-cli 👋</h1>
<p align="center">
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://www.npmjs.com/package/sygnus-cli" target="_blank">
    <img alt="Npm package" src="https://img.shields.io/badge/npm-CB3837?style=for-the-badge&logo=npm&logoColor=white" />
  </a>
</p>

> Cli to generate sygnus modules

## Install

```sh
npm install sygnus-cli --save-dev
```

## 🚀 Usage

Make sure you have [sygnus](https://github.com/Lucasnhso/sygnus) installed

Just run the following command at the root of your project:

```sh
npx sygnus
```

Generate modules:

```sh
npx sygnus generate <module> <name>
```

Or

```sh
npx sygnus g <module> <name>
```

For example:

```sh
#Generate UserRepository
npx sygnus generate repository user

#Generate UserUseCase
npx sygnus generate usecase user

#Generate UserController
npx sygnus generate controller user

#Generate all modules
npx sygnus generate resource user
```

## Author

👤 **Lucas Oliveira**

- Github: [@lucasnhso](https://github.com/lucasnhso)
- LinkedIn: [@lucas-h-oliveira](https://linkedin.com/in/lucas-h-oliveira)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/lucasnhso/sygnus-cli/issues).

## Show your support

Give a ⭐️ if this project helped you!
