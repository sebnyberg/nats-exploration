.PHONY:
re:
	docker-compose restart

.PHONY:
term:
	docker-compose exec -it golang bash
