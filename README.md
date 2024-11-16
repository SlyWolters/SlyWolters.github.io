# Slys Website

## Quick Start Guide

To start the website, run the following command:

```bash
go run main.go
```

> **Note**: Ensure your terminal is in the correct directory. It should look something like:
> `C:/Users/Name/Documents/Programming/slywolters.github.io`.

## Adding Photos Correctly

The file structure for adding photos is as follows:

```bash
static/
    images/
        category1.jpg       # Cover photo for the category
        category1/          # Folder for the category
            item1.jpg       # Main image for the item
            item1/          # Folder for more details about the item
                thumbnail.jpg    # Thumbnail image for the item
                description.txt  # Text description for the item
```

### Details and Instructions

1. **Base Folder**:
   - Place all image files and folders inside `static/images`. This makes them accessible to the website.

2. **Adding a New Category**:
   - Create a new folder within `static/images`. You can name this folder anything you like (e.g., `category1`).
   - Add a cover photo for the category at the top level of `static/images` with the same name as the folder (e.g., `category1.jpg`).

3. **Adding Items to a Category**:
   - Inside the category folder (e.g., `category1`), add a main image (e.g., `item1.jpg`).
   - Create a folder named after the item (e.g., `item1/`) to include additional details.

4. **Adding Thumbnails and Descriptions**:
   - Within the item folder (e.g., `item1/`), include a `thumbnail.jpg` to serve as the preview image.
   - Add a `description.txt` file if you want to provide more information about the item.

> **Example**: For an item named "item1" under "category1":
> - `static/images/category1/item1.jpg` is the main image.
> - `static/images/category1/item1/thumbnail.jpg` is the thumbnail.
> - `static/images/category1/item1/description.txt` holds the item's description.

This structure allows you to organize and display photos and create new categories with ease.