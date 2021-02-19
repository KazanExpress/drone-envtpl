# envtpl plugin for drone

Wraps python [envtpl](https://github.com/andreasjansson/envtpl) as drone plugin.

## Usage

```yaml
- name: injecting env into template
  image: kexpress/drone-envtpl:1.1.0
  settings:
    template: template.j2
    output_file: generated.txt
  environment:
    # template variables
    BUILD_NUMBER: ${DRONE_BUILD_NUMBER}
    TEMPLATE_TEXT: 'Hello world!'
```

template.j2:

```jinja
This file generated using drone envtpl plugin during {{ BUILD_NUMBER }} build.
Here is injected message: {{TEMPLATE_TEXT }}
```

generated.txt

```txt
This file generated using drone envtpl plugin during 1337 build.
Here is injected message: Hello world!
```

## Release Notes

### 1.0

- first release
