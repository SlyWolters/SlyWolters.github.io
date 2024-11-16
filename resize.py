import os
from PIL import Image
import fnmatch

def resize_and_crop_image(img, target_size):
    # Calculate the aspect ratios
    img_ratio = img.width / img.height
    target_ratio = target_size[0] / target_size[1]

    # Resize while maintaining aspect ratio
    if img_ratio > target_ratio:
        # Wider than target: scale based on height
        new_height = target_size[1]
        new_width = int(target_size[1] * img_ratio)
    else:
        # Taller than target: scale based on width
        new_width = target_size[0]
        new_height = int(target_size[0] / img_ratio)

    # Resize the image
    img_resized = img.resize((new_width, new_height), Image.LANCZOS)

    # Calculate cropping box
    left = (new_width - target_size[0]) / 2
    top = (new_height - target_size[1]) / 2
    right = (new_width + target_size[0]) / 2
    bottom = (new_height + target_size[1]) / 2

    # Crop the image to the target size
    img_cropped = img_resized.crop((left, top, right, bottom))
    
    return img_cropped

def resize_images_in_folder(folder_path, output_folder, target_size):
    # Create output folder if it doesn't exist
    os.makedirs(output_folder, exist_ok=True)

    # Traverse the directory and its subdirectories
    for dirpath, _, filenames in os.walk(folder_path):
        for filename in fnmatch.filter(filenames, '*.jpg'):  # Change the pattern as needed
            # Construct full file path
            file_path = os.path.join(dirpath, filename)
            print(f'Resizing and cropping: {file_path}')
            
            # Open an image file
            with Image.open(file_path) as img:
                img_cropped = resize_and_crop_image(img, target_size)
                
                # Construct output path
                relative_path = os.path.relpath(dirpath, folder_path)
                output_path = os.path.join(output_folder, relative_path)

                # Create output subdirectory if it doesn't exist
                os.makedirs(output_path, exist_ok=True)

                # Save the cropped image
                img_cropped.save(os.path.join(output_path, filename))
                
if __name__ == '__main__':
    # Specify the folder containing images
    input_folder = 'static/images'  # Replace with your input folder path
    output_folder = 'images'  # Replace with your output folder path
    new_size = (492, 656)  # Replace with your desired width and height

    resize_images_in_folder(input_folder, output_folder, new_size)
    print('Resizing completed.')
