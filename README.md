# Interests
Web-service for loan interest calculation.

## Install dependencies:
go get -v github.com/gin-gonic/gin

## Build
go build application.go

## Tests
go test ./... -cover -v

## Run
./application

## Test
curl http://127.0.0.1:8080/v1/hello_world/1

{"message":"Hello world"}

## For developers
### SetUp
Для корректной развертки приложения и дальнейшей разработки необходимо клонировать его в по следующему пути
`.../go/src/github.com/apmath-web/`
Клонировать: `git clone https://github.com/apmath-web/template-go`
### Running
Для запуска необходимо запустить билд `application.go`
`go build application.go`
После этого надо запустить полученный бинарник `application`
### Testing
Для прогона всех тестов воспользуйтесь `go test ./... -cover -v` из корневого католога репозитория.
### Работа с Git
Большая просьба не добавлять следующие файлы в репозиторий:
1. /idea
2. application - бинарный файл
3. разлиные настройки вашего окружения/системы
4. кэш фалы тестов и сборок
5. файлы базы данных

Более подробно можно прочесть: [Что не стоит хранить в репозитории](https://clck.ru/FMscw)

Также просьба не использовать 'git push --force' если не знаете что это такое и не удалять чужие ветки, если в них ведется работа.
### Contributing
1. Взять задачу:
Для этого необходимо выбрать одну из задач с label `approve` зайти в нее, прочитать, выбрать себя в `Assignees`, добавить задачу в проект и текущий milestone. После  этого у вас есть время до конца milestone для ее выполнения. Если по каким-то причинам задача не выполнена, то вы можете быть сняты с ее выполнения любым другим человеком. Запрещается брать задачи не помеченные `approve` и задачи для которых не выполнены зависимости.
2. Начать разработку:
После того, как вы выбрали задачу, вы возращаетесь в свой репозиторий и переходите в ветку dev. Отэтой ветки вы создаете ветку под названием feature-`X`, X - номер вашей задачи и работаете в ней. Если на момент начала разработки некоторые зависмости находятся в состоянии UnderReview, то нужно сделать pull из соответствующей ветки.
3. Оформление pull request:
Когда сделан хотя бы один коммит стоит создать pull request(PR). В его офрмление входит то, что вы выбираете свою ветку как
_compare_, а ветку dev как _base_. Так же вы указываете себя в `Assignees`, добавляете PR в проект. После этого в заголовке PR записываете название задачи, а описание переносите список действий из issue с чекбоксами и пишите после этого фразу `close #X`, X - номер задачи. В процессе работы вы помечаете выполненные шаги в чекбоксах.
4. Завершение работы:
По окнчанию работы над задачей вы выбираете в PR label `ready` и выставляете в качестве reviewers @malinink или @levozavr. После этого можно выбирать следующую задачу.
    1. Если PR принят, то все хорошо, можно удалить свою рабочую ветку feature-X
    2. Усли же нет, то в нем будут комметарии к ошибкам и будет снят label `ready`. После исправления ошибок повторно отмечаете reviewer и label `ready`.

### CI-CD
На данный момент в репозитории используется travis. В нем происходит сборка и тестирование кода. Если тесты не проходят, то такой PR не может быть принят. Так же в дальнейшем в нем будет производится и системное тестирование с помощью docker.


## About structure
- Application
  - actions
    - Contain handlers implementation and tests
  - routing
    - Contain router settings
  - view models
    - Contain view models implementation and tests
- Domain
  - Contain interfaces of viewModels, applicationModels, repositories
  - Contain business logic of service
- Infrastructure
  - applicationModels
    - Contain models implementation and tests
  - repositories
    - Contain repositories implementation and tests
- application
  - main package for Application
