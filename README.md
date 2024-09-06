# `extract`

A tool for extracting structured data from unstructured documents. Designed around the Unix philosophy of building simple, modular programs that can be easily connected together to perform more complex tasks.

## Installation

```bash
# TBA
```

## Usage

`extract` supports two modes of operation: `extract` and `infer`. The `extract` mode is used to extract structured data from unstructured documents, while the `infer` mode is used to infer the structure of a document as a JSON schema, which can then be used to extract structured data from similar documents.

### Extract

The `extract` mode is used to extract structured data from unstructured documents. It takes a document and a JSON schema as input, and outputs the structured data extracted from the document.

```bash
extract run --schema <schema> <document>
```

It also supports reading the document from stdin:

```bash
cat <document> | extract run --schema <schema>
curl -s https://somewebsite.com | extract run --schema <schema>
```

And writing the structured data to stdout:

```bash
extract run --schema <schema> <document> > structured_data.json
extract run --schema <schema> <document> | jq ".data"
```

### Infer

The `infer` mode is used to infer the structure of a document as a JSON schema. It takes a document as input, and outputs a JSON schema that describes the structure of the document.

> Note: The `infer` mode is still experimental and may not work as expected. You may need to manually edit the inferred schema to get the desired results.

```bash
extract infer <document>
```

It also supports reading the document from stdin:

```bash
cat <document> | extract infer
curl -s https://somewebsite.com | extract infer
```
