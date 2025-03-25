
1. Install the `opeapi-generator` using homebrew.
```bash
brew install openapi-generator
```

2. Change the BaseURL to `http://localhost:8098` 

3. Generate the swift-package.
```bash
openapi-generator generate \                       
  -i api-definition.yml \
  -g swift5 \
  -o GeneratedMoyaClient \
  --additional-properties=useMoya=true,swiftUseApiNamespace=true
```

4. Add the Swift-Package to the project.
