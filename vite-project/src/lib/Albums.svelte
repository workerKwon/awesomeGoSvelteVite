<script>
    import testapi from "../services/testapi.js";

    let params = {
        id: '',
        title: '',
        artist: '',
        price: ''
    }

    let albums = []

    const getAlbums =  async () => {
        const {data} = await testapi.albums()
        albums = data
    }

    const postAlbum = async () => {
        await testapi.postAlbum(params)
        await getAlbums()
    }

    getAlbums()
</script>

<form on:submit|preventDefault="{postAlbum}">
    <div>id : <input type="text" name="id" bind:value="{params.id}"></div>
    <div>title : <input type="text" name="title" bind:value="{params.title}"></div>
    <div>artist : <input type="text" name="artist" bind:value="{params.artist}"></div>
    <div>price : <input type="number" step="0.01" name="price" bind:value="{params.price}"></div>
    <button type="submit">submit</button>
</form>

{#each albums as album}
    <div>{album.id} | {album.title} | {album.artist} | {album.price} | </div>
{/each}
