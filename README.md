# Terraform Provider for Todoist

This is a custom Terraform provider for managing resources in Todoist, a popular task management application.

## Requirements

- Terraform 0.12.x or later
- Go 1.16.x or later

## Building the provider

Clone the repository:

```bash
$ git clone https://github.com/hilyas/terraform-provider-todoist
```

Enter the provider directory and build the provider:

```bash
$ cd terraform-provider-todoist
$ go build -o terraform-provider-todoist
```

Copy to your Terraform plugins directory for your platform:

```bash
$ cp terraform-provider-todoist ~/.terraform.d/plugins/hilyas.com/hilyas/todoist/0.1.0/darwin_armd64
```

## Using the provider

To use the provider, you must first configure it with your Todoist API key.

```bash
provider "todoist" {
  api_key = "your_todoist_api_key"
}
```

You can also set the `TODOIST_API_KEY` environment variable instead of hardcoding the API key in the configuration:

```bash
$ export TODOIST_API_KEY="your_todoist_api_key"
```

## Resources

- `todoist_project`: This resource represents a Todoist project.

Example usage:

```bash
resource "todoist_project" "example_project" {
  name = "Example Project"
}
```

### Arguments

- `name`: (Required) The name of the project.

## Data Sources

- `todoist_project`: This data source allows you to fetch the details of a Todoist project.

Example usage:

```bash
data "todoist_project" "example_project" {
  project_id = "123456789"
}
```

### Arguments

- `project_id`: (Required) The ID of the project.

## Contributing

If you have suggestions for improvements, bug reports, or other contributions, please open an issue or submit a pull request.

## License

This Terraform provider is released under the MIT License.