# GO-TOKO _API Specification_

## Table of Contents

1. [Introduction](#introduction)
2. [Error Codes](#error-codes)
3. [API Reference](#api-reference)
   1. [Login](#1-login)
   2. [Refresh Token](#2-refresh-token)
   3. [Get All Products](#3-get-all-products)
   4. [Get All Categories](#4-get-all-categories)
   5. [Get All Transactions](#5-get-all-transactions)
   6. [Get Transaction Detail](#6-get-transaction-detail)
   7. [Get Top Products or Categories](#7-get-top-products-or-categories)
   8. [Create Transaction](#8-create-transaction)

## Introduction

API Specification for GO-TOKO Application this document contains the API Specification for GO-TOKO Application to be used by the Frontend Developer and Backend Developer to develop the GO-TOKO Application.

## API Reference

### 1. [Login](api/auth/login.md)

    this endpoint is used to login to the application

### 2. [Refresh Token](api/auth/refresh-token.md)

    this endpoint is used to refresh the token to the application if the token is expired

### 3. [Get All Products](api/products/products.md)

    this endpoint is used to get all products

### 4. [Get All Categories](api/categories/categories.md)

    this endpoint is used to get all categories

### 5. [Get All Transactions](api/transactions/transactions.md)

    this endpoint is used to get all transactions

### 6. [Get Transaction Detail](api/transactions/transaction.md)

    this endpoint is used to get transaction detail

### 7. [Get Top Products or Categories](api/transactions/top.md)

    this endpoint is used to get top products or categories based on transaction

### 8. [Create Transaction](api/transactions/create.md)

    this endpoint is used to create transaction
