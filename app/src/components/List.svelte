<script>
  import Item from './Item.svelte';
  import { onMount } from 'svelte';

  let dwarves = [];

  onMount(async () => {
    const res = await fetch('http://localhost:8080/dwarfs');

    dwarves = await res.json();
  });
</script>

{#each dwarves as dwarf}
  <Item {dwarf} />
{/each}
