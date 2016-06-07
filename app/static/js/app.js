var app = angular.module('Howdy', ['ngWebSocket']);

app.controller("ChatCtrl", ChatCtrl);

function ChatCtrl($scope, $websocket) {
    $scope.responses = ['Messages from your bot will be displayed here.'];
    $scope.sentMessage = 'Please enter your message';

    $scope.ws = $websocket('ws://127.0.0.1:8081/webhook');

    $scope.ws.onMessage(function showResponse(message) {
        var response = JSON.parse(message.data);

        $scope.responses.push(response.text);
        $scope.buttons = [];

        var keyboard = response.reply_markup.keyboard;

        for (var i = 0; i < keyboard.length; i++) {
            for (var j = 0; j < keyboard[i].length; j++) {
                $scope.buttons.push(keyboard[i][j].text);
            }
        }
    });

    $scope.SendMessage = function SendMessage() {
        $scope.sentMessage = $scope.message;

        $scope.validate();

        if (!$scope.error) {
            $scope.send();
        }
    };

    $scope.SendButton = function SendButton(button) {
        $scope.sentMessage = button;

        $scope.validate();

        if (!$scope.error) {
            $scope.send();
        }
    };

    $scope.validate = function validate() {
        $scope.error =
            !$scope.sentMessage ||
            $scope.userId === undefined ||
            !$scope.webhook ||
            !$scope.firstName;
    };

    $scope.send = function send() {
        $scope.message = '';
        $scope.responses = [];

        var payload = JSON.stringify({
            text: $scope.sentMessage,
            userId: $scope.userId,
            webhook: $scope.webhook,
            firstName: $scope.firstName,
            phone: $scope.phone
        });

        $scope.ws.send(payload);
    };
}
