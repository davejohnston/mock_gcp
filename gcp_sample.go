package main

import (
	"log"
	"net/http"

	"encoding/json"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	compute "google.golang.org/api/compute/v1"
)

type Credentials struct {
	Email        string
	PrivateKey   []byte
	PrivateKeyID string
}

func DefaultTokenSource(ctx context.Context, creds *Credentials) (oauth2.TokenSource, error) {

	cfg := &jwt.Config{
		Email:        creds.Email,
		PrivateKey:   creds.PrivateKey,
		PrivateKeyID: creds.PrivateKeyID,
		Scopes:       []string{compute.CloudPlatformScope},
		TokenURL:     "https://accounts.google.com/o/oauth2/token",
	}
	ts := cfg.TokenSource(ctx)

	return ts, nil
}

func CustomClient(ctx context.Context, creds *Credentials) (*http.Client, error) {

	ts, err := DefaultTokenSource(ctx, creds)
	if err != nil {
		return nil, err
	}
	return oauth2.NewClient(ctx, ts), nil

}

var privateKey = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCh1RqAE3kNuEni\nqRV/TKyJj36F4ovVlMjc0WZ+BOt/qAZ+6EPa1KFYOlt3mzUJaU4qd2qYYJ5srAiQ\n8+nhWWJMEnkrwHccj02ID8wLpBc5eDuLSjRol5X/gjH5haiYjC5CEF2jIrN6fCaW\npft9n+dbBM56mnUe4xG/y7m6u3wkBSs2BNG3lO7/5p+NUPl6ctIdo+LvlTlnNu0d\nge7H5jjzH52E/nJ4aD1Lp0HhyHluTlCfR6jXwF3MYLdvuA2WIB63OTq5mBmNh7Oz\nJOLxp57S/4r1uiWCWvW9VvCFX67R4zXCxs2l5CI/8zKb4wnipapvi4yq2M9h8zzA\nVEyYv3ZzAgMBAAECggEARRr8hXuLQlYTSvRcBy+lL9q0MxTiFJLD0TOkGiUcsTb9\nixzAuQX6XfQRkfoXhEzRAaZbwPTlQc+WmttlYILsTELJNmsqEeqJt5PTVJc+CZ7z\ngk/uIEm21lRw5bXzIC0gJbraWWrcjG3Ohp8bTPJG96GOUGF7qPRa5HVo+rnL0Y8r\nIXytgQToWh0ymursU6SGJjKssOu82btZ6UMDQ1KqYgpBGqf80Cgpf0JRuWZicIx4\nSTKYRew5wN1Sm+h8uKaTAPDJqk8DgkJYG9tks1DByzJC+9kLxtNn/wihirsl7MMm\nd1WtjQ3qvJFRAcXGxoWZOTtKKi8ZLdRGpFifyLSsUQKBgQDb5Xj+CehiCFuMIiQU\nJhROWt9qzJJ9CID4GxeOoUPAI5px3uP6fgLOtpq8RFS/AQM1gttqXc5/JrkVO4LQ\nvDJZYRqf6h7BWiAovrzuFvbaHQuQqdTYMlh2lnKEtoWoA2JQtpoEqj19I/hjRX/3\nHZq9219jazIAq7MRl1Uig8sfIwKBgQC8ZyD4Jef/WlUOq4TwbfhY3HNw32keJ0id\nVCZlH8d42GfkblpAI86/LrGO8+aBKw1JuJd2Tc/vYuigpGSnL6z2mu8qEe38t8DR\n5/vO+gtVSfXqVTnaIbwnP6iA8d3uYiZJJAMjG4kCN92dW9dIrtHayKEXWNd2+4oe\nNepPrhXocQKBgEuC1HagHl2zswQ/IRaOMQDrMbdyAofPKMKKQx+mfyZ2021w9eap\n6PrlN+OYr2VbyqdQhVpUVjnRqVMCOZzGU5/fuY3ajq8k7NAxF53G4wPpX3RQ7ZdE\nSp6GcVLjfqhAaT2ARwl9EFptxLkKP7QzRVUXBP2V7PjP/VD4H7MXgOPzAoGBAJb9\no7ucTYklsSNXrOwvhihZTR95vToQS67jIP5McMXV0bWOB0B+MhSgbGbf6607fqPF\nj4WdqPb2cu9DsPMYT2s4ElLKGcw+zAat/+4KEQkihDZwZTP/c1aVOwtUTAPfn0Vg\n2i7Jw40GjtKtJyU2DjNkNU6Lweq2fyPlG8sN/rrRAoGALkC5ylIHiZGTN2w1ObDg\nHPCle55W73Af5YskyRo9Wj5xhLqVOyqfb0pazjNb8FFKTvqGxBBObFuYLd9EKTat\n70uSi8HebHuKllzwPb+ZXGzo8AcmDpY72ge8w8IS9snpVzsfLGdcXE5Y8pK21mw8\nfcDdhnJ2gVCT9nc6c14m2Oo=\n-----END PRIVATE KEY-----\n"

func main() {
	ctx := context.Background()
	c, err := CustomClient(ctx, &Credentials{
		Email:        "bob@valid-tuner-195415.iam.gserviceaccount.com",
		PrivateKey:   []byte(privateKey),
		PrivateKeyID: "abcee3dd28fbbd3bf8c47b21094aa8f81d72833e"})

	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	// Project ID for this request.
	project := "valid-tuner-195415" // TODO: Update placeholder value.

	req := computeService.Instances.List(project, "us-east1-b")
	if err := req.Pages(ctx, func(page *compute.InstanceList) error {
		for _, instance := range page.Items {
			//log.Printf("Zone [%s] Instance: %s", "us-east1-b", instance.Name)
			data, _ := json.MarshalIndent(instance, "", "\t")
			log.Print(string(data))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

}
