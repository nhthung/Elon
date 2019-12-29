import torch
import torch.nn as nn
import torch.optim as optim
import torchvision.transforms as transforms
import torchvision.datasets as datasets


INPUT_SIZE = 28 * 28
OUTPUT_SIZE = 10
NUM_EPOCHS = 30
LEARNING_RATE = 3.0


class Net(nn.Module):
    def __init__(self):
        super(Net, self).__init__()
        self.h1 = nn.Linear(INPUT_SIZE, 100)
        self.out = nn.Linear(100, OUTPUT_SIZE)

    def forward(self, x):
        x = torch.sigmoid(self.h1(x))
        x = torch.sigmoid(self.out(x))
        return x


def run_network(net):
    mse_loss = nn.MSELoss()
    sgd = torch.optim.SGD(net.paramters(), lr=LEARNING_RATE)

    train_loader = get_train_loader()
    train_network(
        train_loader, net, NUM_EPOCHS, sgd,
        create_input_reshaper(),
        create_loss_function(mse_loss)
        )
    print()

    test_loader = get_test_loader()
    test_network(
        test_loader, net,
        create_input_reshaper()
        )
    