## apigeecli registry apis versions artifacts get

Get Artifact details for an API version

### Synopsis

Get Artifact details for an API version

```
apigeecli registry apis versions artifacts get [flags]
```

### Options

```
      --api-name string      API Name
      --api-version string   API Version
  -c, --content              If set to true, returns artifact contents
  -h, --help                 help for get
  -n, --name string          Artifact Name
```

### Options inherited from parent commands

```
  -a, --account string     Path Service Account private key in JSON
      --default-token      Use Google default application credentials access token
      --disable-check      Disable check for newer versions
      --metadata-token     Metadata OAuth2 access token
      --no-output          Disable printing all statements to stdout
      --no-warnings        Disable printing warnings to stderr
  -o, --org string         Apigee organization name
      --print-output       Control printing of info log statements (default true)
  -p, --projectID string   Apigee Registry Project ID; Use if Apigee Orgniazation not provisioned in this project
      --region string      Region where Apigee Registry is provisioned
  -t, --token string       Google OAuth Token
```

### SEE ALSO

* [apigeecli registry apis versions artifacts](apigeecli_registry_apis_versions_artifacts.md)	 - Manage artifacts for an API version

###### Auto generated by spf13/cobra on 6-Mar-2024
