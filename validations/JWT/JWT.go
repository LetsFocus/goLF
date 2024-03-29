package JWT

import (
	"crypto/rsa"
	"github.com/LetsFocus/goLF/errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type Keys struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func CreateJWTToken(algorithm, subject string, header, data map[string]interface{}, keys Keys) (string, error) {
    switch algorithm {
    case "RSA":
        token := jwt.New(jwt.SigningMethodRS512)

        // Set token headers
        for key, value := range header {
            token.Header[key] = value
        }

        claims := token.Claims.(jwt.MapClaims)

        // Set token claims
        for key, value := range data {
            claims[key] = value
        }

        // Set subject and expiration claims
        claims["sub"] = subject
        claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

        // Sign the token
        tokenString, err := token.SignedString(keys.PrivateKey)
        if err != nil {
            return "", err
        }

        return tokenString, nil
    }

    return "", nil
}

func GetRSAPrivateKey(privateKeyString string) (*rsa.PrivateKey, error) {
	pemPrivateKey := "-----BEGIN PRIVATE KEY-----\n" + privateKeyString + "\n-----END PRIVATE KEY-----"
	decodedPrivateKey := []byte(pemPrivateKey)

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return nil, &errors.Errors{StatusCode: http.StatusBadRequest,
			Code: http.StatusText(http.StatusBadRequest), Reason: "invalid PrivateKey " + err.Error()}
	}

	return key, nil
}

func ValidateJWTToken(refreshToken string, publicKey *rsa.PublicKey) (map[string]string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return "", &errors.Errors{StatusCode: http.StatusBadRequest,
				Code: http.StatusText(http.StatusBadRequest), Reason: "refreshToken is invalid"}
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, &errors.Errors{StatusCode: http.StatusBadRequest,
			Code: http.StatusText(http.StatusBadRequest), Reason: "refresh token is invalid " + err.Error()}
	}

	claimsData := make(map[string]string)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var (
			subOk, audOk, issOK bool
		)

		claimsData["sub"], subOk = claims["sub"].(string)
		claimsData["aud"], audOk = claims["aud"].(string)
		claimsData["iss"], issOK = claims["iss"].(string)

		if !subOk || !audOk || !issOK {
			return nil, &errors.Errors{StatusCode: http.StatusBadRequest,
				Code: http.StatusText(http.StatusBadRequest), Reason: "data is missing in the refresh token"}
		}
	}

	return claimsData, nil
}

func GetRSAPublicKey(publicKeyString string) (*rsa.PublicKey, error) {
	pemPublicKey := "-----BEGIN PUBLIC KEY-----\n" + publicKeyString + "\n-----END PUBLIC KEY-----"
	decodedPublicKey := []byte(pemPublicKey)

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, &errors.Errors{StatusCode: http.StatusBadRequest,
			Code: http.StatusText(http.StatusBadRequest), Reason: "invalid PublicKey " + err.Error()}
	}
	
	return key, nil
}

func generateKeyPair() (*rsa.PrivateKey, error) {
    // Generate a new RSA key pair
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, err
    }

    return privateKey, nil
}

func savePrivateKeyToPEM(privateKey *rsa.PrivateKey, filename string) error {
    // Serialize the private key to PEM format
    privateKeyPEM := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
    }

    // Write the private key to a file
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    err = pem.Encode(file, privateKeyPEM)
    if err != nil {
        return err
    }

    return nil
}

func savePublicKeyToPEM(publicKey *rsa.PublicKey, filename string) error {
    // Serialize the public key to PEM format
    publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        return err
    }

    publicKeyPEM := &pem.Block{
        Type:  "RSA PUBLIC KEY",
        Bytes: publicKeyBytes,
    }

    // Write the public key to a file
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    err = pem.Encode(file, publicKeyPEM)
    if err != nil {
        return err
    }

    return nil
}
