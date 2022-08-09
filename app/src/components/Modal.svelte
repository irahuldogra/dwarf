<script lang="ts">
  import { closeModal } from 'svelte-modals';
  export let title: string;
  export let isOpen: boolean;
  export let redirect: string;
  export let dwarf: string;
  export let random: boolean;
  export let send: Function;

  let data = {
    redirect: redirect,
    dwarf: dwarf,
    random: random,
  };
</script>

{#if isOpen}
  <div role="dialog" class="modal">
    <div class="contents">
      <h2>{title}</h2>
      <label for="redirect">Redirect To:</label>

      <input
        type="text"
        name="redirect"
        id="redirect"
        bind:value={data.redirect}
      />

      <label for="dwarf">Dwarf</label>
      <input
        type="text"
        name="dwarf"
        id="dwarf"
        bind:value={data.dwarf}
        disabled={data.random}
        class={data.random ? 'disabled' : ''}
      />

      <div class="isRandom">
        <label for="isRandom"> Randomly Generated Dwarf </label>
        <input
          class="checkbox"
          type="checkbox"
          name="isRandom"
          id="isRandom"
          bind:checked={data.random}
        />
      </div>

      <div class="actions">
        <button on:click={closeModal}>Cancel</button>
        <button
          on:click={() => {
            send(data);
            closeModal();
          }}>Send</button
        >
      </div>
    </div>
  </div>
{/if}

<style>
  .modal {
    position: fixed;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    z-index: 1;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: rgb(0, 0, 0);
    background-color: rgba(0, 0, 0, 0.2);
    -webkit-transition: 0.5s;
    overflow: auto;
    transition: all 0.3s linear;
    /* allow click-through to backdrop */
    pointer-events: none;
  }

  .contents {
    min-width: 400px;
    padding: 20px;
    color: black;
    background-color: white;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    pointer-events: auto;
    border: 3px solid;
    box-shadow: 5px 5px 0px 0px, 2px 2px 0px 0px, 4px 4px 0px 0px,
      6px 6px 0px 0px, 8px 8px 0px 0px;
    user-select: none;
    -webkit-user-select: none;
    touch-action: manipulation;
  }

  .contents h2 {
    font-size: 1.7em;
    border-left: 6px solid black;
    padding: 10px;
    color: #000;
    letter-spacing: 4px;
    margin-bottom: 40px;
    font-weight: bold;
    padding-left: 10px;
  }

  h2 {
    text-align: center;
    font-size: 24px;
  }

  .actions {
    margin-top: 32px;
    display: flex;
    justify-content: space-between;
    gap: 8px;
  }

  .disabled {
    background: slategray;
  }

  button {
    font-family: 'Open Sans', sans-serif;
    font-size: 16px;
    letter-spacing: 2px;
    text-decoration: none;
    text-transform: uppercase;
    color: #000;
    cursor: pointer;
    border: 3px solid;
    padding: 0.25em 0.5em;
    box-shadow: 1px 1px 0px 0px, 2px 2px 0px 0px, 3px 3px 0px 0px,
      4px 4px 0px 0px, 5px 5px 0px 0px;
    position: relative;
    user-select: none;
    -webkit-user-select: none;
    touch-action: manipulation;
  }

  button:active {
    box-shadow: 0px 0px 0px 0px;
    top: 5px;
    left: 5px;
  }

  @media (min-width: 768px) {
    button {
      padding: 0.25em 0.75em;
    }
  }

  label {
    font-size: 16px;
    letter-spacing: 1.2px;
    margin-left: 12px;
  }

  input[type='text'] {
    font-size: 0.9em;
    font-weight: 600;
    padding: 0.25em 0.8em;
    margin-bottom: 10px;
    line-height: 2em;
    border: 3px solid black;
    border-left: 5px solid black;
    border-radius: 20px;
  }

  input[type='text']:disabled {
    background-color: slategray;
    color: lightgray;
  }

  .isRandom {
    display: grid;
    grid-template-columns: 15em auto;
    gap: 1em;
    align-items: center;
  }

  input[type='checkbox'].checkbox {
    font-size: 15px;
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    width: 3.5em;
    height: 1.56em;
    background: lightgray;
    border-radius: 3em;
    position: relative;
    cursor: pointer;
    outline: none;
    -webkit-transition: all 0.2s ease-in-out;
    transition: all 0.2s ease-in-out;
  }

  input[type='checkbox'].checkbox:checked {
    background: black;
  }

  input[type='checkbox'].checkbox:after {
    position: absolute;
    content: '';
    width: 1.5em;
    height: 1.5em;
    border-radius: 50%;
    background: #fff;
    -webkit-box-shadow: 0 0 0.25em rgba(0, 0, 0, 0.3);
    box-shadow: 0 0 0.25em rgba(0, 0, 0, 0.3);
    -webkit-transform: scale(0.7);
    transform: scale(0.7);
    left: 0;
    -webkit-transition: all 0.2s ease-in-out;
    transition: all 0.2s ease-in-out;
  }

  input[type='checkbox'].checkbox:checked:after {
    left: calc(100% - 1.5em);
  }
</style>
