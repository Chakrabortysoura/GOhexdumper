
# GO HexDumper

A hexDumper written in GO with a modern look made with the help of lipgloss styling library.

Supports all the basic features of the major fetures of the hexDumper utility in most linux systems but with a modern nicer look. Improvments in how the utility represents when displaying the same data in various different representation such as -

a. Hexadecimal b. Decimal c. Hexadecimal in Canonical form(Big-endian form representation)

(Also support for length and offset parameter)




## Demo

s![Screenshot_20250610_220400](https://github.com/user-attachments/assets/6aaeffa8-28aa-4f75-b9c5-266e13a83d7e)


This is a testcase use shown in the above image to demonstrate different color coding for various data formats.

## Input Parameters Manual

#### Show the manual for all the input parameters

```http
--help
```

#### Decimal Representation

```http
  -d
```

#### Octal Representation

```http
  -o
```

#### Hexademical Canonical Form Representation

```http
  -c
```

#### Binary Representation

```http
  -b
```
#### Total Byte Length Read

```http
  -n [int]
```
#### Start Offset Byte Length

```http
  -s [int]
```
