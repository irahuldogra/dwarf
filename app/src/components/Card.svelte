<div class="card">
  <slot />
</div>

<style>
  .card {
    background-color: whitesmoke;
    color: black;
    font-weight: 400;
    letter-spacing: 0.12em;
    border-radius: 15px;
    padding: 40px;
    margin: 20px 10px;
    max-width: 800px;
    box-shadow: rgba(255, 255, 255, 0.2) 0px 0px 0px 1px inset,
      rgba(30, 30, 30, 0.9) 0px 0px 0px 1px;
  }
</style>
