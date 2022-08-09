<script lang="ts">
  import Card from './Card.svelte';
  import Modal from './Modal.svelte';
  import { Modals, closeModal, openModal } from 'svelte-modals';
  export let dwarf: Dwarf;

  let showCard = true;
  interface Dwarf {
    id: number;
    dwarf: string;
    redirect: string;
    random: boolean;
    clicked: boolean;
  }

  async function update(data: Dwarf) {
    const json = {
      redirect: data.redirect,
      dwarf: data.dwarf,
      random: data.random,
      id: dwarf.id,
    };

    await fetch('http://localhost:8080/dwarfs', {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(json),
    }).then((response) => {
      console.log(response);
    });
  }

  function handleOpen(dwarf: Dwarf) {
    openModal(Modal, {
      title: 'Update Dwarf Link',
      send: update,
      dwarf: dwarf.dwarf,
      redirect: dwarf.redirect,
      random: dwarf.random,
    });
  }

  async function deleteDwarf() {
    if (
      confirm(
        'Are your sure you want to delete this dwarf link (' +
          dwarf.dwarf +
          ')?'
      )
    ) {
      await fetch(`http://localhost:8080/dwarfs/${dwarf.id}`, {
        method: 'DELETE',
      }).then((response) => {
        showCard = false;
        console.log(response);
      });
    }
  }
</script>

{#if showCard}
  <Card>
    <p>Dwarf : http://localhost:8080/r/{dwarf.dwarf}</p>
    <p>Redirect : {dwarf.redirect}</p>
    <p>Clicked : {dwarf.clicked}</p>
    <button class="update" on:click={() => handleOpen(dwarf)}> Update</button>
    <button class="delete" on:click={deleteDwarf}> Delete</button>
  </Card>

  <Modals>
    <div slot="backdrop" class="backdrop" on:click={closeModal} />
  </Modals>
{/if}

<style>
  .update {
    background-color: yellowgreen;
  }
  .delete {
    background-color: red;
  }
  .backdrop {
    position: fixed;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    background: rgb(255, 255, 255);
  }

  p {
    box-shadow: 0px 4px 0px rgb(30, 30, 30);
  }

  button {
    font-family: 'Open Sans', sans-serif;
    font-size: 16px;
    margin-right: 0.3rem;
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
</style>
