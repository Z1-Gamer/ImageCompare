<script>
  import { getContext } from "svelte"
  export var message
  const { close } = getContext("simple-modal")

  function CloseThis() {
    close()
  }
</script>

<div class="Popup">
  <p>
    {message}
  </p>
  <button id="CloseButton" on:click={CloseThis}> OK </button>
</div>

<style>
  .Popup {
    align-self: center;
    text-align: center;
    width: 260px;
    height: 100px;
  }
  #CloseButton {
    background-color: #79e8df;
    border-radius: 6px;
    border-color: #79e8df;
    text-shadow: none;
  }
</style>
