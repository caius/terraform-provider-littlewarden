# terraform-provider-littlewarden

Terraform provider for <https://littlewarden.com>'s API, to allow managing your site checks.

## Usage

Setup the provider with your Little Warden API key (found in the sidebar of <https://littlewarden.com/api>)

```terraform
provider "littlewarden" {
  api_key = "secret-key-here"
}
```

## License

See `LICENSE` for details.
