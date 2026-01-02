class Things3Cli < Formula
  desc "CLI for Things 3"
  homepage "https://github.com/ossianhempel/things3-cli"
  url "https://github.com/ossianhempel/things3-cli/archive/0bc1a5fdaf935636ae7b68f5e35c45615010dc35.tar.gz"
  sha256 "c3c18210d4fe6bcc9e4db446b90382452fd7d211a01a8b6b94fcd9b5b4e32d34"
  version "0bc1a5f"

  depends_on "go" => :build

  def install
    ldflags = "-s -w -X github.com/ossianhempel/things3-cli/internal/cli.Version=#{version}"
    system "go", "build", *std_go_args(ldflags: ldflags), "./cmd/things"
  end

  test do
    system "#{bin}/things", "--version"
  end
end
