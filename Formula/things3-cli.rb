class Things3Cli < Formula
  desc "CLI for Things 3"
  homepage "https://github.com/ossianhempel/things3-cli"
  url "https://github.com/ossianhempel/things3-cli/archive/32f13fecfc6bccdafbfa1a461b716545f1fc4917.tar.gz"
  sha256 "ef9b7b50197f2142b9dbe7e231fafae04b751a8545584a822dcf3c7ff48965b9"
  version "32f13fe"

  depends_on "go" => :build

  def install
    ldflags = "-s -w -X github.com/ossianhempel/things3-cli/internal/cli.Version=#{version}"
    system "go", "build", "-trimpath", "-ldflags", ldflags, "-o", bin/"things", "./cmd/things"
  end

  test do
    system "#{bin}/things", "--version"
  end
end
