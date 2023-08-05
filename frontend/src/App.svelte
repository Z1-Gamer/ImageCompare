<script lang="ts">
  import { CompareFiles, DebugMe } from "../wailsjs/go/main/App.js"
  import EmptyImage from "./assets/images/BlankImageT2.png"
  import EmptyResultImage from "./assets/images/BlankImageT3.png"
  import Modal from "svelte-simple-modal"
  import PopupManager from "./PopupManager.svelte"

  let ContentComponent: PopupManager
  let error: string

  // Component Binded Variables
  var canvas1: HTMLCanvasElement
  var canvas2: HTMLCanvasElement
  var canvas3: HTMLCanvasElement
  var CompareButton: HTMLButtonElement
  var ResultLabel: HTMLParagraphElement

  // Window Size Variables
  var WindowWidth = 750
  var WindowHeight = 900
  var PopupWidth = 300
  var PopupHeight = 100

  // Image Variables
  const AllowedImageTypes = ["image/png", "image/jpeg"]
  var CurrentImage1: ImageBitmap
  var CurrentImage2: ImageBitmap
  var Image1Type: string
  var Image2Type: string
  var Image1Base64: string
  var Image2Base64: string
  var CompareResultImageBase64: string
  var CompareResultPercent: Number

  const eventLog = document.querySelector(".event-log-contents")

  // Get the Default Image and draw it on canvas1 and canvas2.
  const defaultimage = new Image()
  defaultimage.src = EmptyImage
  defaultimage.onload = function () {
    canvas1.width = 250
    canvas1.height = 250
    const ctx1 = canvas1.getContext("2d")
    ctx1.shadowColor = "Blue"
    ctx1.shadowBlur = 0
    ctx1.drawImage(defaultimage, 5, 5, 240, 240)

    // Set the current image variable for canvas 1 to the default image
    createImageBitmap(defaultimage).then((response) => {
      CurrentImage1 = response
    })

    canvas2.width = 250
    canvas2.height = 250
    const ctx2 = canvas2.getContext("2d")
    ctx2.shadowColor = "Blue"
    ctx2.shadowBlur = 0
    ctx2.drawImage(defaultimage, 5, 5, 240, 240)

    // Set the current image variable for canvas 2 to the default image
    createImageBitmap(defaultimage).then((response) => {
      CurrentImage2 = response
    })
  }

  //Get the Default Result Image and draw it on canvas 3.
  const DefaultResultImage = new Image()
  DefaultResultImage.src = EmptyResultImage
  DefaultResultImage.onload = function () {
    canvas3.width = 240
    canvas3.height = 240
    const ctx3 = canvas3.getContext("2d")
    ctx3.drawImage(DefaultResultImage, 0, 0, 240, 240)
  }

  // Create a dragoever event listener for the whole document to:
  // Disable file drop if the target is not either Canvas1 or Canvas2, and if the dragged items' type is not in the AloowedImageTypes.
  document.addEventListener("dragover", function (event) {
    event.preventDefault()
    let dtarget = event.target as Element
    if ((dtarget.id != "Canvas1" && dtarget.id != "Canvas2") || !AllowedImageTypes.includes(event.dataTransfer.items[0].type)) {
      event.dataTransfer.dropEffect = "none"
    }
  })

  // Adjust the size of components when resizing the window
  window.onresize = function () {
    if (window.innerWidth != WindowWidth) {
      canvas1.style.width = window.innerWidth / 3 + "px"
      canvas2.style.width = window.innerWidth / 3 + "px"
      canvas3.style.width = window.innerWidth / 3 + "px"
      CompareButton.style.top =
        parseInt(window.getComputedStyle(canvas1, null).getPropertyValue("top")) +
        -(parseInt(window.getComputedStyle(canvas1, null).getPropertyValue("height")) * 1.5) / 3 +
        "px"
      CompareButton.style.width = window.innerWidth / 7.5 + "px"
      CompareButton.style.height = window.innerWidth / 25 + "px"

      WindowWidth = window.innerWidth
    }
    var WindowHeight = window.innerHeight
  }

  // When entering either canvas 1 or 2 while dragging a file add a highlighting effect using canvas shadowblur.
  function handleDragEnter(ev: DragEvent, canvasnum: Number) {
    switch (canvasnum) {
      case 1: {
        canvas1.getContext("2d").clearRect(0, 0, canvas1.width, canvas1.height)
        if (CurrentImage1.width < CurrentImage1.height) {
          let ImageRatio = CurrentImage1.height / CurrentImage1.width
          canvas1.height = 240 * ImageRatio + 10
        } else {
          let ImageRatio = CurrentImage1.width / CurrentImage1.height
          canvas1.height = 240 / ImageRatio + 10
        }

        canvas1.getContext("2d").shadowColor = "Blue"
        canvas1.getContext("2d").shadowBlur = 20
        canvas1.getContext("2d").drawImage(CurrentImage1, 5, 5, canvas1.width - 10, canvas1.height - 10)
        break
      }
      case 2: {
        canvas2.getContext("2d").clearRect(0, 0, canvas2.width, canvas2.height)
        if (CurrentImage2.width < CurrentImage2.height) {
          let ImageRatio = CurrentImage2.height / CurrentImage2.width
          canvas2.height = 240 * ImageRatio + 10
        } else {
          let ImageRatio = CurrentImage2.width / CurrentImage2.height
          canvas2.height = 240 / ImageRatio + 10
        }

        canvas2.getContext("2d").shadowColor = "Blue"
        canvas2.getContext("2d").shadowBlur = 20
        canvas2.getContext("2d").drawImage(CurrentImage2, 5, 5, canvas2.width - 10, canvas2.height - 10)
        break
      }
    }
  }

  function handleDragLeave(ev: DragEvent, canvasnum: Number) {
    switch (canvasnum) {
      case 1: {
        canvas1.getContext("2d").clearRect(0, 0, canvas1.width, canvas1.height)
        if (CurrentImage1.width < CurrentImage1.height) {
          let ImageRatio = CurrentImage1.height / CurrentImage1.width
          canvas1.height = 240 * ImageRatio + 10
        } else {
          let ImageRatio = CurrentImage1.width / CurrentImage1.height
          canvas1.height = 240 / ImageRatio + 10
        }

        canvas1.getContext("2d").shadowBlur = 0
        canvas1.getContext("2d").drawImage(CurrentImage1, 5, 5, canvas1.width - 10, canvas1.height - 10)
        break
      }
      case 2: {
        canvas2.getContext("2d").clearRect(0, 0, canvas2.width, canvas2.height)
        if (CurrentImage2.width < CurrentImage2.height) {
          let ImageRatio = CurrentImage2.height / CurrentImage2.width
          canvas2.height = 240 * ImageRatio + 10
        } else {
          let ImageRatio = CurrentImage2.width / CurrentImage2.height
          canvas2.height = 240 / ImageRatio + 10
        }

        canvas2.getContext("2d").shadowBlur = 0
        canvas2.getContext("2d").drawImage(CurrentImage2, 5, 5, canvas2.width - 10, canvas2.height - 10)
        break
      }
    }
  }

  // When dropping a file: first check to see it is an image file. Then use a reader to create a Base64 string of it and store it in a variable.
  // Create an ImageBitmap from the file, store it in the CurrentImage(1/2) variable and draw it on the canvas over which it was dropped.
  function handleDrop(ev: DragEvent, canvasnum: Number) {
    ev.preventDefault()

    if (ev.dataTransfer.items[0].type.includes("image")) {
      const file = ev.dataTransfer.items[0].getAsFile()

      let read = new FileReader()
      read.readAsDataURL(file)
      read.onloadend = function () {
        let filebuf = read.result as string
        if (canvasnum == 1) {
          Image1Base64 = filebuf
          Image1Type = file.type
          console.log(Image1Type)
        }
        if (canvasnum == 2) {
          Image2Base64 = filebuf
          Image2Type = file.type
          console.log(Image2Type)
        }
        // For Debugging
        // DebugMe(filebuf)

        createImageBitmap(file).then((response) => {
          const canvas = document.getElementById("Canvas" + canvasnum) as HTMLCanvasElement

          if (canvasnum == 1) {
            CurrentImage1 = response
          }
          if (canvasnum == 2) {
            CurrentImage2 = response
          }
          if (response.width < response.height) {
            let ImageRatio = response.height / response.width
            canvas.height = 240 * ImageRatio + 10
          } else {
            let ImageRatio = response.width / response.height
            canvas.height = 240 / ImageRatio + 10
          }

          const ctx = canvas.getContext("2d")
          ctx.clearRect(0, 0, canvas.width, canvas.height)
          ctx.drawImage(response, 5, 5, canvas.width - 10, canvas.height - 10)

          // Center the button vertically depending on the canvas height. Note: Surely there must be an easier way but both TS and HTML coding are new to me so there. :)
          CompareButton.style.top =
            parseInt(window.getComputedStyle(canvas, null).getPropertyValue("top")) +
            -(parseInt(window.getComputedStyle(canvas, null).getPropertyValue("height")) * 1.5) / 3 +
            "px"
        })
      }
    }
  }

  /*
  Call the CompareImage function in go (app.go) unless there are no images added to both canvas 1 and 2 or if they have different sizes, then show an error popup. 
  Otherwise:
  After receiveng the resulting image and percent, first convert the base64 to an actual image and then make an ImageBitmap.
  Draw the first image (canvas1) with alpha 0.6 and draw the RedGreen result on top of it to get a nice effect.
  Show the resulting percentage on the label 
  */
  function CompareImages() {
    if (Image1Base64 != null && Image2Base64 != null) {
      if (CurrentImage1.width == CurrentImage2.width && CurrentImage1.height == CurrentImage2.height) {
        CompareFiles(Image1Base64, Image2Base64, Image1Type, Image2Type).then((result) => {
          CompareResultImageBase64 = result.ImageData
          CompareResultPercent = result.Percent
          fetch(CompareResultImageBase64)
            .then((response) => response.blob())
            .then((blob) => {
              const ImageResultFile = new File([blob], "result.png", {
                type: "image/png",
              })
              createImageBitmap(ImageResultFile).then((response) => {
                canvas3.getContext("2d").clearRect(0, 0, canvas3.width, canvas3.height)
                canvas3.getContext("2d").globalAlpha = 0.5
                if (CurrentImage1.width < CurrentImage1.height) {
                  let ImageRatio = CurrentImage1.height / CurrentImage1.width
                  canvas3.height = 240 * ImageRatio
                } else {
                  let ImageRatio = CurrentImage1.width / CurrentImage1.height
                  canvas3.height = 240 / ImageRatio
                }
                canvas3.getContext("2d").drawImage(CurrentImage1, 0, 0, canvas3.width, canvas3.height)
                canvas3.getContext("2d").globalAlpha = 0.6
                canvas3.getContext("2d").drawImage(response, 0, 0, canvas3.width, canvas3.height)
                ResultLabel.textContent = "The images are " + CompareResultPercent.toFixed(0) + "% identical."
              })
            })
        })
      } else {
        error = "Error: Images have different sizes!"
        ContentComponent.ShowPopup(error, PopupWidth, PopupHeight)
      }
    } else {
      error = "Error: Please add some images first!"
      ContentComponent.ShowPopup(error, PopupWidth, PopupHeight)
    }
  }
</script>

<main>
  <p id="Title">Image Compare</p>
  <div id="MainDiv">
    <canvas
      bind:this={canvas1}
      id="Canvas1"
      height="240"
      width="240"
      on:dragenter={(event) => handleDragEnter(event, 1)}
      on:dragleave={(event) => handleDragLeave(event, 1)}
      on:drop={(event) => handleDrop(event, 1)}
    />

    <button bind:this={CompareButton} id="Button1" on:click={CompareImages}> Compare </button>

    <canvas
      bind:this={canvas2}
      id="Canvas2"
      height="240"
      width="240"
      on:dragenter={(event) => handleDragEnter(event, 2)}
      on:dragleave={(event) => handleDragLeave(event, 2)}
      on:drop={(event) => handleDrop(event, 2)}
    />
  </div>

  <p bind:this={ResultLabel} id="ResultLabel">You will see the difference as percentage here</p>
  <canvas bind:this={canvas3} id="Canvas3" height="240" width="240" />
</main>

<Modal><PopupManager bind:this={ContentComponent} /></Modal>

<style>
  #Title {
    font-family: sans-serif;
    margin: auto;
    padding-top: 1%;
  }
  #Canvas1 {
    position: relative;
    left: -5%;
    padding-top: 2%;
  }
  #Canvas2 {
    position: relative;
    left: +5%;
    padding-top: 2%;
  }
  #Canvas3 {
    position: relative;
    top: -10px;
    align-self: center;
  }
  #Button1 {
    position: relative;
    font-family: sans-serif;
  }
  #ResultLabel {
    font-size: 15px;
    font-family: sans-serif;
    position: relative;
    top: -10px;
  }
  :global(body) {
    font-family: sans-serif;
    background-color: rgb(20, 30, 55);
  }
</style>
