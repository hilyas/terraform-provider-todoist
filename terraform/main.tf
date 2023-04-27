terraform {
  required_providers {
    todoist = {
      source = "hilyas.com/hilyas/todoist"
      version = "0.1.0"
    }
  }
}

provider "todoist" {
  api_key = "9e098c0362e2d5ee63576b64e0eeb3a94a21eeef"
}

resource "todoist_project" "example_project" {
  name = "Example Project"
}

output "project_id" {
  value = todoist_project.example_project.id
}

output "project_name" {
  value = todoist_project.example_project.name
}

data "todoist_project" "example_project_data" {
  project_id = todoist_project.example_project.id
}

output "fetched_project_id" {
  value = data.todoist_project.example_project_data.id
}

output "fetched_project_name" {
  value = data.todoist_project.example_project_data.name
}
