# WordPress Plugin Generator

A WordPress plugin generator that creates PSR-4 compliant boilerplate code, making it easier to start new WordPress plugin projects.

## Features

- Generates PSR-4 compliant WordPress plugin structure
- Includes basic plugin boilerplate files
- Sets up automated class autoloading
- Generates plugin header with standard WordPress metadata
- Includes WordPress stubs

## Installation

```bash
go install github.com/damian-elenbaas/wp-plugin-generator@latest
```

## Usage

```bash
wp-plugin-generator generate CompanyName/ProjectName --export-dir ~/directory/of/your/project
```

### Options

- `--export-dir`: The directory where you want the plugin to be exported

## Generated Structure

The generated structure is based on [this boilerplate](https://github.com/DevinVinson/WordPress-Plugin-Boilerplate).

```
my-awesome-plugin/
├── admin/
│   ├── css/
│   ├── js/
│   └── ProjectName_Admin.php
├── api/
│   └── ProjectName_API.php
├── includes/
│   ├── ProjectName_Loader.php
│   └── (and more...)
├── public/
│   ├── css/
│   ├── js/
│   └── ProjectName_Public.php
├── composer.json
├── index.php
└── my-awesome-plugin.php
```

## Requirements

- Go 1.24 or higher
- PHP 8.4 or higher (for generated plugins)
- Composer (for managing generated plugin dependencies)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
