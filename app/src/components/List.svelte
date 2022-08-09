<script>
  import Item from './Item.svelte';
  import { onMount } from 'svelte';
  import Constant from '../constants';

  let dwarves = [];

  onMount(async () => {
    const res = await fetch(`${Constant.SERVICE_URL}/dwarfs`);

    dwarves = await res.json();
  });
</script>

{#each dwarves as dwarf}
  <Item {dwarf} />
{/each}
