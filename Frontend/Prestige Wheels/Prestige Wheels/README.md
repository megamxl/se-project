
1. Install the `opeapi-generator` using homebrew.
```bash
brew install openapi-generator
```

2. `cd` into `se-project/Rental-Server`

3. Generate the swift-package.
```bash
openapi-generator generate \
  -i api-definition.yml \
  -g swift5 \
  -o GeneratedMoyaClient \
  --additional-properties=swiftUseApiNamespace=true
```

4. move the folder `GeneratedMoyaClient` to `se-project/Frontend`

5. Add the Swift-Package to the project.
