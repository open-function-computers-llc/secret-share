<h1>Hello! Please feel free to share a secret:</h1>
<form method="POST" action="/store">
    <textarea name="secret" cols="30" rows="10"></textarea>
    <label for="viewCount">
        <span>How many times should this secret be viewable?</span>
        <input type="number" min="3" max="10" name="viewCount" value="5" />
    </label>
    <label for="autoshare">
        <span>Autoshare with Open Function Computers, LLC</span>
        <input type="checkbox" name="autoshare" id="autoshare">
    </label>
    <label>
        <input type="submit" value="Share" />
    </label>
</form>
