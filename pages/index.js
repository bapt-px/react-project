import Link from 'next/link'
import React, { useState } from 'react';

export default () => {

  const [url, setUrl] = useState('')
  const [nbComment, setNbComment] = useState(0)
  const [average, setAverage] = useState(0)

  const searchPage = e => {
      let formData = new URLSearchParams()
      formData.append('url', url)
    fetch('http://www.localhost:8083', {
      method: 'POST',
      mode: 'cors',
      headers: new Headers(),
      body: formData
    })
      .then(response => response.json())
        .then(data => {
        setAverage(data.average)
        setNbComment(data.nbComment)
    })
  }

    const divStyle = {
        textAlign: "center"
    }

  return (
    <div style={divStyle}>
        <h1>RateMyTrip</h1>
      <input type="text" name="url" value={url} onChange={e => setUrl(e.target.value)} /><br/>
      <button onClick={() => searchPage()}>
        Click me
      </button>
      <p>
        Words: {nbComment}<br/>
        Average:: {average}
      </p>

  </div>
  )
}
