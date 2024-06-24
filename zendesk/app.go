package zendesk

import (
       "fmt"
       "io/ioutil"
)

func (c *client) GetAppPublicKey(appID int64) (string, error) {
       endpoint := fmt.Sprintf("/api/v2/apps/%d/public_key.pem", appID)
       resp, err := c.request("GET", endpoint, nil, nil)
       if err != nil {
               return "", err
       }
       defer resp.Body.Close()
       cert, err := ioutil.ReadAll(resp.Body)
       if err != nil {
               return "", err
       }

       return string(cert), err
}
