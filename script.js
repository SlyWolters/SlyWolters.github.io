function zoomImage(event) {
    const container = document.querySelector('.zoom-container');
    const image = document.getElementById('zoomImage');

    const containerRect = container.getBoundingClientRect();
    const mouseX = event.clientX - containerRect.left;
    const mouseY = event.clientY - containerRect.top;

    const percentX = mouseX / containerRect.width * 100;
    const percentY = mouseY / containerRect.height * 100;

    const translateX = (percentX - 50) * 2; // Adjust the multiplier for the desired zoom level
    const translateY = (percentY - 50) * 2;

    image.style.transform = `translate(${translateX}%, ${translateY}%) scale(2)`;
  }