<script lang="ts">
  import Button from "./atoms/Button.svelte";
  import LoginDropDown from "./LoginDropDown.svelte";

  let commentInput: HTMLElement
  
  let onInput = () => {
      commentInput.style.height = commentInput.scrollHeight / 2 + 'px'
      commentInput.style.height = commentInput.scrollHeight + 'px'
    }
</script>

  <div>
    <textarea on:input="{onInput}" bind:this="{commentInput}" name="comment" rows="3" placeholder="Write your comment"
    class="overflow-hidden shadow-sm px-3 py-3 block w-full text-base font-normal focus:outline-none border border-slate-200 resize-none rounded-md"
    ></textarea>
    <div class="flex flex-row">
      <input class="basis-2/3 shadow-sm px-3 py-3 my-2 mr-2 text-base font-normal focus:outline-none border border-slate-200 rounded-md"
              type="text" name="" id="" placeholder="Name">
      <div class="basis-1/3 m-2">
        <Button value={"Comment as Guest"}/>
      </div>
    </div>
      <LoginDropDown/>
  </div>

<style>
</style>
