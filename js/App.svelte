<script lang="ts">
import { each } from "svelte/internal";
 let photos = [""];
  fetch("/api/photos").then((resp) => {
    return resp.json() as Promise<string[]>
  }).then((p) => {
    photos = p
  })
</script>

<main>
{#each photos as p}
<img src="/images/{p}" alt="Photo {p}">
{/each}
</main>

<style>
	img {
      width: 30%;
      padding: 1%;
  }
</style>
