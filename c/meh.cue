apiVersion: "traefik.containo.us/v1alpha1"
kind:       "TLSStore"
metadata: {
	name:      "traefik"
	namespace: "traefik"
}

spec: defaultCertificate: secretName: "defn-run-wildcard"
