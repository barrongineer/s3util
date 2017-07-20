// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"io/ioutil"

	"bytes"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Upsert a resource in S3",
	Long: `Example:
	s3util put --key path/to/something/cool.exe --file ./cool.exe`,
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := viper.GetString("endpoint")
		bucket := viper.GetString("bucket")
		key := viper.GetString("key")
		filename := viper.GetString("file")
		id := viper.GetString("id")
		secret := viper.GetString("secret")
		region := viper.GetString("region")

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		contentType := http.DetectContentType(data)
		body := bytes.NewReader(data)
		size := int64(len(data))
		encryption := "aws:kms"
		input := &s3.PutObjectInput{
			Bucket:               &bucket,
			Key:                  &key,
			Body:                 body,
			ContentLength:        &size,
			ContentType:          &contentType,
			ServerSideEncryption: &encryption,
		}

		creds := credentials.NewStaticCredentials(id, secret, "")
		cfg := aws.NewConfig().
			WithRegion(region).
			WithCredentials(creds).
			WithEndpoint(endpoint)
		s := session.Must(session.NewSession())
		s3Client := s3.New(s, cfg)
		_, err = s3Client.PutObject(input)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(putCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	putCmd.PersistentFlags().StringP("key", "k", "", "key to use for the file")
	viper.BindPFlag("key", putCmd.PersistentFlags().Lookup("key"))

	putCmd.PersistentFlags().StringP("file", "f", "", "local path to the file")
	viper.BindPFlag("file", putCmd.PersistentFlags().Lookup("file"))
}
