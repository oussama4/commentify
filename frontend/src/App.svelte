<script lang="ts">
  import { onMount } from 'svelte';

  import CommentArea from './components/CommentArea.svelte'
  import Thread from './components/Thread.svelte';
  import { Comment as CommentType } from "./interfaces/comment";
    
  let comments: CommentType[];

  onMount(async () => {
    const commentsEndpoint = `${import.meta.env.VITE_COMMENTIFY_BASE_URL}/comments?page_url=https://iaeineq.net/rfziprfuqdna`;
    const res = await fetch(commentsEndpoint, {
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      }
    });
    const resBody = await res.json();
    comments = resBody.comments.map((c) => {
      return {
        body: c.body,
        author: {
          name: c.author.name
        }
      }
    });
  });
</script>

<main class="md:w-1/3 mx-auto mt-5">

  <CommentArea />
  <Thread comments={comments}/>

</main>

<style>
</style>
