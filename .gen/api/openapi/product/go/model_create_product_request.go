/*
 * Sweep API
 *
 * Manage Product
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type CreateProductRequest struct {

	Sku string `json:"sku"`

	Name string `json:"name"`

	Expirable bool `json:"expirable"`
}
