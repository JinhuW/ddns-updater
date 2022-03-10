# Cloudflare

Before setting this in the ddns-updater, Plase creare a dns record in the cloudflare.

<img width="1111" alt="image" src="https://user-images.githubusercontent.com/13999458/157681573-7b7da77b-7c4d-4c15-a0e3-d091dc7ab85d.png">

The setting here only update the dns record instead of creating a dns record.

## Configuration

### Example

```json
{
  "settings": [
    {
      "provider": "cloudflare",
      "zone_identifier": "some id",
      "identifier": "<IDENTIFIER>",
      "domain": "domain.com",
      "host": "@",
      "ttl": 600,
      "token": "yourtoken",
      "ip_version": "ipv4"
    }
  ]
}
```

### Compulsory parameters

- `"zone_identifier"` is the Zone ID of your site
- "identifier" is the ID of your API Token (Step 4 in Domain setup)
- `"domain"`
- `"host"` is your host. It should be left to `"@"`, since subdomain and wildcards (`"*"`) are not really supported by Cloudflare it seems.
See [this issue comment for context](https://github.com/qdm12/ddns-updater/issues/243#issuecomment-928313949). This is left as is for compatibility.
- `"ttl"` integer value for record TTL in seconds (specify 1 for automatic)
- One of the following:
    - Email `"email"` and Global API Key `"key"`
    - User service key `"user_service_key"`
    - API Token `"token"`, configured with DNS edit permissions for your DNS name's zone

### Optional parameters

- `"proxied"` can be set to `true` to use the proxy services of Cloudflare
- `"ip_version"` can be `ipv4` (A records) or `ipv6` (AAAA records), and defaults to `ipv4 or ipv6`

## Domain setup

1. Make sure you have `curl` installed
1. Obtain your API key from Cloudflare website ([see this](https://support.cloudflare.com/hc/en-us/articles/200167836-Where-do-I-find-my-Cloudflare-API-key-))
1. Obtain your zone identifier for your domain name, from the domain's overview page written as *Zone ID*
1. Find your **identifier** in the `id` field with

    ```sh
    ZONEID=aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
    EMAIL=example@example.com
    APIKEY=aaaaaaaaaaaaaaaaaa
    curl -X GET "https://api.cloudflare.com/client/v4/zones/$ZONEID/dns_records" \
        -H "X-Auth-Email: $EMAIL" \
        -H "X-Auth-Key: $APIKEY"
    ```
    
    ```json
    {
        "result":[
            {
                "id":"sfdhjalsdjfhlasjhdfa",          <------ IDENTIFIER 
                "zone_id":"1231231231dsfasdf",
                "zone_name":"******",
                "name":"******",
                "type":"A",
                "content":"108.49.66.116",
                "proxiable":true,
                "proxied":true,
                "ttl":1,
                "locked":false,
                "meta":{
                    "auto_added":false,
                    "managed_by_apps":false,
                    "managed_by_argo_tunnel":false,
                    "source":"primary"
                },
                "created_on":"2022-03-10T05:07:01.397266Z",
                "modified_on":"2022-03-10T05:07:01.397266Z"
            }
        ],
        "success":true,
        "errors":[
            
        ],
        "messages":[
            
        ],
        "result_info":{
            "page":1,
            "per_page":100,
            "count":1,
            "total_count":1,
            "total_pages":1
        }
    }
    ```

You can now fill in the necessary parameters in *config.json*

Special thanks to @Starttoaster for helping out with the [documentation](https://gist.github.com/Starttoaster/07d568c2a99ad7631dd776688c988326) and testing.
