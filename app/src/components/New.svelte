<script>
  import { openModal } from 'svelte-modals';
  import Modal from './Modal.svelte';

  async function updateDwarf(data) {
    const json = {
      redirect: data.redirect,
      dwarf: data.dwarf,
      random: data.random,
    };
    await fetch('http://localhost:8080/dwarfs', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(json),
    }).then((response) => {
      console.log(response);
    });
  }

  function handleOpen() {
    openModal(Modal, {
      title: 'Create New Dwarf Link',
      send: updateDwarf,
      redirect: '',
      dwarf: '',
      random: false,
    });
  }
</script>

<button on:click={handleOpen}>New</button>

<style>
  button {
    background-color: darkgray;
    font-family: 'Open Sans', sans-serif;
    font-size: 16px;
    letter-spacing: 2px;
    text-decoration: none;
    text-transform: uppercase;
    color: whitesmoke;
    cursor: pointer;
    border: 3px solid;
    padding: 0.25em 0.5em;
    box-shadow: 1px 1px 0px 0px, 2px 2px 0px 0px, 3px 3px 0px 0px,
      4px 4px 0px 0px, 5px 5px 0px 0px;
    touch-action: manipulation;
    user-select: none;
    -webkit-user-select: none;
    -ms-transform: skewX(-5deg);
    -webkit-transform: skewX(-5deg);
    transform: skewX(-5deg);
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
</style>
