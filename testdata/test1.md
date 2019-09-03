	---
	title: "Architecture"
	date: 2019-09-01T14:47:42+02:00
	draft: true
	author: doublemalt
	---
	
	Deployment
	----------
	
	
	```plantuml
	@startuml
	actor Author
	actor Editor
	actor Reviewer
	
	cloud "Ethereum Blockchain" {
	  component "Peer Review Contract" as prc
	  component "Journal Contract" as jc
	  agent "Author Wallet" as aw
	  agent "Editor Wallet" as ew
	  agent "Reviewer Wallet" as rw
	}
	
	
	cloud "IPFS Network" {
	  database "Smart Contract UI Code" {
		folder "Peer Review Contract UI Code" as prc_ui
		folder "Journal UI Code" as jc_ui
	  }
	  database "Review Artifacts" {
		folder "Submitted Paper Versions" as spv
		folder "Paper Reviews" as pr
	  }
	}
	
	Author == aw
	aw -(0- prc
	'Author -- prc_ui
	Author -->> spv : uploads
	Author <<-- pr : reads
	
	Editor == ew
	ew -(0- jc
	'Editor -- jc_ui
	ew -(0- prc
	'Editor -- prc_ui
	Editor <<-- spv : reads
	Editor <<-- pr : reads
	
	Reviewer == rw
	rw -(0- prc
	Reviewer <<-- spv : reads
	'Reviewer -- prc_ui
	Reviewer -->> pr : uploads
	
	prc -(0)- jc
	@enduml
	```

	more text