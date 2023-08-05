<script>
  import { getContext } from "svelte"
  import Popup from "./Popup.svelte"
  const { open, close } = getContext("simple-modal")
  export const ShowPopup = (error = new String(), width = new Number(), height = new Number()) =>
    open(
      Popup,
      { message: error },
      {
        closeButton: false,
        closeOnOuterClick: false,
        styleBg: {
          background: "rgba(0, 0, 0, 0.5)",
        },
        styleWindowWrap: {
          margin: "4rem",
        },
        styleWindow: {
          width: width + "px",
          height: height + "px",
          border: "3px solid rgb(20, 100, 220)",
          boxShadow: "inset 0 0 0 2px white, 0 0 0 2px white",
          background: "#ff8000",
        },
      }
    )
</script>
