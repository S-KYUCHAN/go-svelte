<script>
    import axios from "axios";
    import { onMount } from 'svelte';

    const article = {
        title: "",
        description: "",
        content: ""
    };

    async function postAPI() {
        await axios.post(
            `http://localhost:8000/article`,
            JSON.stringify(article)
        ).then(response => console.log(response));
    }

    const handleSubmit = () => {
        console.log(article);
        postAPI();
    };
</script>

<form on:submit|preventDefault={handleSubmit}>
    <input type="text" placeholder= "Title" bind:value={article.title} />
    <input type="text" placeholder= "Description" bind:value={article.description} />
    <input type="text" placeholder= "Content" bind:value={article.content} />
    <button>Add Article</button>
</form>

<h1> Title is {article.title}</h1>

<h2> description : {article.description}</h2>

<h3> content </h3>

{article.content}

