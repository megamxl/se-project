
Install the `opeapi-generator` using homebrew.
```zsh
brew install openapi-generator
```

Generate the swift-package.
```zsh
openapi-generator generate \                       
  -i api-definition.yml \
  -g swift5 \
  -o GeneratedMoyaClient \
  --additional-properties=useMoya=true,swiftUseApiNamespace=true
```

Add the Swift-Package to the project.
