<script lang="ts">
    export let name: String
    const colors = ['bg-slate-800', 'bg-red-800', 'bg-yellow-800', 'bg-lime-800', 'bg-sky-800', 'bg-blue-800', 'bg-violet-800']

    let randomNumberFromText = (text) => {
        return text.split('')
            .map(char => char.charCodeAt(0))
            .reduce((current, next) => current + next)
    }
    let randomColor = () => {
        return colors[randomNumberFromText(name) % colors.length]
    }
</script>

<div class="{randomColor()} inline-block h-16 w-16 rounded-full ring-2 ring-white text-center text-white text-5xl"> {name[0]} </div>